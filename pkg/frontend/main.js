'use strict';
let oamap = L.map('oamap').fitWorld().setView([53.773056, 20.476111], 12);
L.tileLayer.wms('http://msipmo.olsztyn.eu/arcgis/services/msipmo_Plan/MapServer/WMSServer?', {
    layers: '0,1,2,3,4,5,6,7,8,9,10,11,12',
    tileSize: 256,
    format: 'image/png',
    maxZoom: 16,
}).addTo(oamap);

function onLocationFound(e) {
    L.circle(e.latlng, (e.accuracy / 2)).addTo(oamap);
}

oamap.on('locationfound', onLocationFound);
oamap.locate({setView: false, maxZoom: 16});

let availableRoutes;
let vehiclesLayerGroups = [];

function InternalVehicle(rawVehicle) {
    this.stall = (rawVehicle.route.length === 0);

    if (this.stall) {
        this.route = rawVehicle.next_route;
        this.trip_id = rawVehicle.next_trip_id;
        this.description = rawVehicle.next_description;
        this.vector = 0;
    } else {
        this.route = rawVehicle.route;
        this.trip_id = rawVehicle.trip_id;
        this.description = rawVehicle.description;
        this.vector = 180 + rawVehicle.vector;
    }

    this.latitude = rawVehicle.latitude;
    this.longitude = rawVehicle.longitude;
    this.variance = rawVehicle.variance;

    this.isStall = function () {
        return this.stall;
    };
}

function serializeVehicles(rawVehicles) {
    let serialized = [];
    rawVehicles.forEach(v => {
        serialized.push(new InternalVehicle(v));
    });
    return serialized;
}

function getVehicleIcon(internalVehicle) {
    let iconHtml = '<svg width="26" height="26">\n' +
        '<path transform="rotate({{ROTATE}} 13 19)" fill="{{FILL}}" stroke="{{STROKE}}" stroke-width="2"\n' +
        'stroke-dasharray="{{STROKE-DASHARRAY}}" d="\n' +
        'M26\n' +
        '19c0-2.2-0.6-4.4-1.6-6.2C22.2\n' +
        '8.8\n' +
        '13\n' +
        '0\n' +
        '13\n' +
        '0S3.8\n' +
        '8.7\n' +
        '1.6\n' +
        '12.8c-1\n' +
        '1.8-1.6\n' +
        '4-1.6\n' +
        '6.2c0\n' +
        '7.2\n' +
        '5.8\n' +
        '13\n' +
        '13\n' +
        '13\n' +
        'S26\n' +
        '26.2\n' +
        '26\n' +
        '19 Z"/>\n' +
        '<g>\n' +
        '<text x="13" y="19" font-family="sans-serif" font-size="12px" fill="white"\n' +
        'text-anchor="middle" alignment-baseline="central">{{ROUTE}}\n' +
        '</text>\n' +
        '</g>\n' +
        'Sorry, your browser does not support inline SVG.\n' +
        '</svg>';

    let getRouteColor = function (determinedRoute) {
        switch (determinedRoute) {
            case 'N01':
            case 'N02':
                return 'rgba(0, 0, 0, 0.9)';
            case '1':
            case '2':
            case '3':
                return 'rgba(227, 30, 30, 0.9)';
        }
        return 'rgba(0, 157, 210, 0.9)';
    };

    let getIconBorderColor = function (variance) {
        if (variance > -120) {
            return 'white';
        } else if (variance > -600) {
            return 'yellow';
        }
        return 'orange';
    };

    let getIconBorderStyle = function (internalVehicle) {
        if (internalVehicle.isStall()) {
            return '2,2';
        }
        return '0,0';
    };

    iconHtml = iconHtml.replace("{{ROTATE}}", internalVehicle.vector);
    iconHtml = iconHtml.replace("{{FILL}}", getRouteColor(internalVehicle.route));
    iconHtml = iconHtml.replace("{{STROKE}}", getIconBorderColor(internalVehicle.variance));
    iconHtml = iconHtml.replace("{{STROKE-DASHARRAY}}", getIconBorderStyle(internalVehicle));
    iconHtml = iconHtml.replace("{{ROUTE}}", internalVehicle.route);

    return iconHtml;
}

function markerizeVehicles(internalVehicles) {
    let markerized = [];
    internalVehicles.forEach(v => {
        let opts = {
            icon: L.divIcon({
                html: getVehicleIcon(v),
                className: "outerVehicleIcon",
            }),
            alt: v.route,
            riseOnHover: true,
            title: v.description,
        };

        let popupContent =
            "<b>Linia nr " + v.route + "</b><br>" +
            "Numer kursu: " + v.trip_id + "<br>" +
            "Kierunek: " + v.description + "<br>";
        if (v.isStall()) {
            popupContent = popupContent.concat("Czas do odjazdu: " + v.variance + "s");
        } else {
            popupContent = popupContent.concat("Opóźnienie: " + (-1 * v.variance) + "s");
        }

        markerized.push(L.marker([v.latitude, v.longitude], opts).bindPopup(popupContent));
    });
    return markerized;
}

function insertOnMap(rawVehicles) {
    let serializedVehicles = serializeVehicles(rawVehicles);
    let markerizedVehicles = markerizeVehicles(serializedVehicles);

    availableRoutes.forEach(r => {
        vehiclesLayerGroups[r.route].clearLayers();
        for (let i = 0; i < serializedVehicles.length; i++) {
            if (serializedVehicles[i].route === r.route) {
                vehiclesLayerGroups[r.route].addLayer(markerizedVehicles[i]);
            }
        }
    });
}


function fireNextRefresh(lastModifiedDate) {
    let refreshAfter = 22000 - (Date.now() - lastModifiedDate);
    if (refreshAfter < 0) {
        if (refreshAfter > -21000) {
            setTimeout(refreshMap, 1000);
        } else {
            console.log("Refreshing the map has been stalled. " +
                "This can happen when there's a problem with the data source. " +
                "Reload the page to start the refreshing again.")
        }
    } else {
        setTimeout(refreshMap, refreshAfter);
    }
}

function refreshMap() {
    fetch('http://localhost:8080/VehiclesData')
        .then(function (response) {
            let lastModified = new Date(response.headers.get("Last-Modified"));
            fireNextRefresh(lastModified);
            return response.json();
        })
        .then(function (responseJSON) {
            insertOnMap(JSON.parse(JSON.stringify(responseJSON)));
        });
}

function initializeMap() {
    fetch('http://localhost:8080/AvailableRoutes')
        .then(function (response) {
            return response.json();
        })
        .then(function (responseJSON) {
            availableRoutes = JSON.parse(JSON.stringify(responseJSON));

            availableRoutes.forEach(r => {
                vehiclesLayerGroups[r.route] = new L.LayerGroup();
            });
            L.control.layers(null, vehiclesLayerGroups).addTo(oamap).expand();

            refreshMap();
        });
}

initializeMap();