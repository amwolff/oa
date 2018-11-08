'use strict';
let availableRoutes;
let vehiclesLayerGroups = [];

function InternalVehicle(rawVehicle) {
    this.stall = (rawVehicle.route.length === 0);

    if (this.stall) {
        this.route = rawVehicle.next_route;
        this.trip_id = rawVehicle.next_trip_id;
        this.description = rawVehicle.next_description;
        this.azimuth = 0;
    } else {
        this.route = rawVehicle.route;
        this.trip_id = rawVehicle.trip_id;
        this.description = rawVehicle.description;
        this.azimuth = 180 + rawVehicle.azimuth;
    }

    this.latitude = rawVehicle.latitude;
    this.longitude = rawVehicle.longitude;
    this.variance = rawVehicle.variance;

    this.isStall = function () {
        return this.stall;
    };

    this.getMarker = function () {
        // TODO(amwolff): that would be better than markerizing already internal data type
    }
}

function serializeVehicles(rawVehicles) {
    let serialized = [];
    rawVehicles.forEach(v => {
        serialized.push(new InternalVehicle(v));
    });
    return serialized;
}

function getVehicleIcon(internalVehicle) {
    let iconTemplate = '<svg class="outerVehicleIcon" width="26" height="26">\n' +
        '<path transform="rotate({rotate} 13 19)" fill="{fill}" stroke="{stroke}" stroke-width="2"\n' +
        'stroke-dasharray="{stroke_dasharray}" d="\n' +
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
        'text-anchor="middle" alignment-baseline="central">{route}\n' +
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

    let iconData = {
        rotate: internalVehicle.azimuth,
        fill: getRouteColor(internalVehicle.route),
        stroke: getIconBorderColor(internalVehicle.variance),
        stroke_dasharray: getIconBorderStyle(internalVehicle),
        route: internalVehicle.route,
    };

    return L.Util.template(iconTemplate, iconData);
}

// TODO(amwolff): make betterSeconds more informative name
function betterSeconds(seconds) {
    if (seconds > 60) {
        let minutes = Math.floor(seconds / 60);
        let remainingSeconds = seconds - 60 * minutes;
        if (minutes > 60) {
            let hours = Math.floor(minutes / 60);
            let remainingMinutes = minutes - 60 * hours;
            return L.Util.template('{HH}h {MM}m {SS}s', {
                HH: hours,
                MM: remainingMinutes,
                SS: remainingSeconds,
            });
        }
        return L.Util.template('{MM}m {SS}s', {MM: minutes, SS: remainingSeconds});
    }
    return L.Util.template('{SS}s', {SS: seconds});
}

function markerizeVehicles(internalVehicles) {
    let markerized = [];
    internalVehicles.forEach(v => {
        let opts = {
            icon: L.divIcon({
                html: getVehicleIcon(v),
                className: 'innerVehicleIcon',
            }),
            alt: v.route,
            riseOnHover: true,
            title: v.description,
        };

        let popupTemplate =
            '<b>Linia nr {route}</b><br>' +
            'Numer kursu: {trip_id}<br>' +
            'Kierunek: {description}<br>' +
            '{state}: {variance}';

        let popupData = {
            route: v.route,
            trip_id: v.trip_id,
            description: v.description,
        };

        if (v.isStall()) {
            if (v.variance > 0) {
                popupData.state = 'Czas do odjazdu';
                popupData.variance = betterSeconds(v.variance);
            } else {
                popupData.state = 'Opóźnienie odjazdu';
                popupData.variance = betterSeconds(-1 * v.variance);
            }
        } else {
            if (v.variance > 0) {
                popupData.state = 'Przed czasem';
                popupData.variance = betterSeconds(v.variance);
            } else {
                popupData.state = 'Opóźnienie';
                popupData.variance = betterSeconds(-1 * v.variance);
            }
        }

        markerized.push(L.marker([v.latitude, v.longitude], opts).bindPopup(L.Util.template(popupTemplate, popupData)));
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

// TODO(amwolff): swap elses with returns
function fireNextRefresh(lastModifiedDate) {
    let refreshAfter = 22000 - (Date.now() - lastModifiedDate);
    if (refreshAfter < 0) {
        if (refreshAfter > -21000) {
            setTimeout(refresh, 1000);
        } else {
            console.log(
                "Refreshing the map has been stalled. " +
                "This can happen when there's a problem with the data source. " +
                "Reload the page to start the refreshing again.");
        }
    } else {
        setTimeout(refresh, refreshAfter);
    }
}

function refresh() {
    fetch('https://api.autobusy.olsztyn.pl/Vehicles')
        .then(function (response) {
            let lastModified = new Date(response.headers.get('Last-Modified'));
            fireNextRefresh(lastModified);
            return response.json();
        })
        .then(function (responseJSON) {
            insertOnMap(JSON.parse(JSON.stringify(responseJSON)));
        });
}

function entrypoint() {
    let map = L.map('map', {attributionControl: false, center: [53.773056, 20.476111], zoom: 14});

    // TODO(amwolff): can we afford adding the {r} (retina tiles) parameter to the tile layer URL?
    L.tileLayer('https://api.mapbox.com/styles/v1/amwolff/cjnynkofj1jxf2ro9v4123t0v/tiles/256/{z}/{x}/{y}?access_token={t}', {
        t: 'pk.eyJ1IjoiYW13b2xmZiIsImEiOiJjamtndGVqMnUwbjV2M3BueDRxNWtqODQ5In0.f6Sd2mM-5ozz45F4ZxlU8Q',
        minZoom: 9,
        maxZoom: 18,
    }).addTo(map);

    L.control.attribution({
        prefix: false,
        position: 'bottomright',
    }).addAttribution(
        '<a href="mailto:kontakt@autobusy.olsztyn.pl">Zgłoś błąd</a>' +
        ' / © <a href="https://www.mapbox.com/about/maps/">Mapbox</a>' +
        ' © <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>' +
        ' <strong><a href="https://www.mapbox.com/map-feedback/" target="_blank">Improve this map</a></strong>').addTo(map);

    let userLocation;

    let onLocationFound = function (e) {
        let r = e.accuracy / 2;
        if (map.hasLayer(userLocation)) {
            userLocation.setLatLng(e.latlng).setRadius(r);
            return;
        }
        userLocation = L.circle(e.latlng, {
            radius: r,
            color: '#FF6C00',
        }).addTo(map);
        map.flyToBounds(userLocation.getBounds(), {maxZoom: 17});
    };

    let onLocationError = function (e) {
        console.log(e.message);
    };

    map.on('locationfound', onLocationFound);
    map.on('locationerror', onLocationError);
    map.locate({watch: true, enableHighAccuracy: true});

    fetch('https://api.autobusy.olsztyn.pl/Routes')
        .then(function (response) {
            return response.json();
        })
        .then(function (responseJSON) {
            availableRoutes = JSON.parse(JSON.stringify(responseJSON));

            availableRoutes.forEach(r => {
                vehiclesLayerGroups[r.route] = new L.LayerGroup();
            });
            L.control.layers(null, vehiclesLayerGroups).addTo(map).expand();

            refresh();
        });
}

entrypoint();