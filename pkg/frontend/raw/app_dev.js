'use strict';
const iconTemplate =
    '<svg class="outerVehicleIcon" width="26" height="26">\n' +
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

const popupTemplate =
    '<b>Linia nr {route}</b><br>' +
    'Numer kursu: {trip_id}<br>' +
    'Kierunek: {description}<br>' +
    '{state}: {variance}';

class InternalVehicle {
    constructor(rawVehicle) {
        this.stall = (rawVehicle.route.length === 0);

        if (this.stall) {
            this.azimuth = 0;
            this.description = rawVehicle.next_description;
            this.route = rawVehicle.next_route;
            this.trip_id = rawVehicle.next_trip_id;
        } else {
            this.azimuth = 180 + rawVehicle.azimuth;
            this.description = rawVehicle.description;
            this.route = rawVehicle.route;
            this.trip_id = rawVehicle.trip_id;
        }

        this.latitude = rawVehicle.latitude;
        this.longitude = rawVehicle.longitude;
        this.variance = rawVehicle.variance;
    };

    static getIconBodyColor(route) {
        switch (route) {
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

    static getIconBorderColor(variance) {
        if (variance > -120) {
            return 'white';
        } else if (variance > -600) {
            return 'yellow';
        }
        return 'orange';
    };

    static getIconBorderStyle(stall) {
        if (stall) {
            return '2,2';
        }
        return '0,0';
    };

    static hrSeconds(seconds) {
        if (seconds > 60) {
            const minutes = Math.floor(seconds / 60);
            const remainingSeconds = seconds - 60 * minutes;
            if (minutes > 60) {
                const hours = Math.floor(minutes / 60);
                const remainingMinutes = minutes - 60 * hours;
                return L.Util.template('{HH}h {MM}m {SS}s', {
                    HH: hours,
                    MM: remainingMinutes,
                    SS: remainingSeconds,
                });
            }
            return L.Util.template('{MM}m {SS}s', { MM: minutes, SS: remainingSeconds });
        }
        return L.Util.template('{SS}s', { SS: seconds });
    };

    getIconHTML() {
        const iconData = {
            rotate: this.azimuth,
            fill: InternalVehicle.getIconBodyColor(this.route),
            stroke: InternalVehicle.getIconBorderColor(this.variance),
            stroke_dasharray: InternalVehicle.getIconBorderStyle(this.stall),
            route: this.route,
        };

        return L.Util.template(iconTemplate, iconData);
    };

    getPopupContent() {
        const popupData = {
            route: this.route,
            trip_id: this.trip_id,
            description: this.description,
        };

        if (this.stall) {
            if (this.variance > 0) {
                popupData.state = 'Czas do odjazdu';
                popupData.variance = InternalVehicle.hrSeconds(this.variance);
            } else {
                popupData.state = 'Opóźnienie odjazdu';
                popupData.variance = InternalVehicle.hrSeconds(-1 * this.variance);
            }
        } else {
            if (this.variance > 0) {
                popupData.state = 'Przed czasem';
                popupData.variance = InternalVehicle.hrSeconds(this.variance);
            } else {
                popupData.state = 'Opóźnienie';
                popupData.variance = InternalVehicle.hrSeconds(-1 * this.variance);
            }
        }

        return L.Util.template(popupTemplate, popupData);
    };

    getLeafletDivIcon() {
        return L.divIcon({
            html: this.getIconHTML(),
            className: 'innerVehicleIcon',
        });
    };

    getLeafletLatLng() {
        return L.latLng(this.latitude, this.longitude);
    };

    getLeafletMarker() {
        const opts = {
            icon: this.getLeafletDivIcon(),
            title: this.route,
            alt: this.trip_id,
            riseOnHover: true,
        };

        return L.marker(this.getLeafletLatLng(), opts).bindPopup(this.getPopupContent(), { autoPan: false });
    };
}

function serializeVehicles(rawVehicles) {
    const serialized = [];
    rawVehicles.forEach(v => {
        serialized.push(new InternalVehicle(v));
    });
    return serialized;
}

function updateMarker(dstMarker, srcInternalVehicle) {
    dstMarker.setIcon(srcInternalVehicle.getLeafletDivIcon());
    dstMarker.setLatLng(srcInternalVehicle.getLeafletLatLng());

    const popup = dstMarker.getPopup();
    popup._content = srcInternalVehicle.getPopupContent();
    popup.update();
}

let availableRoutes;

const vehiclesLayerGroups = [];

function insertOnMap(rawVehicles) {
    const serializedVehicles = serializeVehicles(rawVehicles);

    availableRoutes.forEach(r => {
        vehiclesLayerGroups[r.route].eachLayer(o => {
            let found = false;
            serializedVehicles.forEach(n => {
                if (n.trip_id === o.options.alt) {
                    found = true;
                }
            });
            if (!found) {
                vehiclesLayerGroups[r.route].removeLayer(o);
            }
        });

        serializedVehicles.forEach(n => {
            if (n.route === r.route) {
                let found = false;
                vehiclesLayerGroups[r.route].eachLayer(o => {
                    if (o.options.alt === n.trip_id) {
                        found = true;
                        updateMarker(o, n);
                    }
                });
                if (!found) {
                    vehiclesLayerGroups[r.route].addLayer(n.getLeafletMarker());
                }
            }
        });
    });
}

// function insertOnMap(rawVehicles) {
//     const serializedVehicles = serializeVehicles(rawVehicles);
//
//     availableRoutes.forEach(r => {
//         vehiclesLayerGroups[r.route].clearLayers();
//         serializedVehicles.forEach(v => {
//             if (v.route === r.route) {
//                 vehiclesLayerGroups[r.route].addLayer(v.marker);
//             }
//         });
//     });
// }

// TODO(amwolff): swap elses with returns
function planNextRefresh(lastModifiedDate) {
    const refreshAfter = 7000 - (Date.now() - lastModifiedDate);
    if (refreshAfter < 0) {
        if (refreshAfter > -6000) {
            setTimeout(refresh, 1000);
        } else {
            console.log(
                'Refreshing the map has been stalled. ' +
                'This can happen when there\'s a problem with the data source or with your clock. ' +
                'Reload the page to start the refreshing again.');
            // TODO(amwolfff): Maybe 'setTimeout(refresh, 12000);'
        }
    } else {
        setTimeout(refresh, refreshAfter);
    }
}

const endpointVehicles = 'http://localhost:8080/Vehicles';

function refresh() {
    fetch(endpointVehicles)
        .then(response => {
            const lastModified = new Date(response.headers.get('Last-Modified'));
            planNextRefresh(lastModified);
            return response.json();
        })
        .then(responseJSON => {
            insertOnMap(JSON.parse(JSON.stringify(responseJSON)));
        });
}

function setLocationTracking(map) {
    let userLocation;

    const onLocationFound = function (e) {
        const r = e.accuracy / 2;
        if (map.hasLayer(userLocation)) {
            userLocation.setLatLng(e.latlng).setRadius(r);
            return;
        }
        userLocation = L.circle(e.latlng, {
            radius: r,
            color: '#FF6C00',
        }).addTo(map);
        map.flyToBounds(userLocation.getBounds(), { maxZoom: 17 });
    };

    const onLocationError = function (e) {
        console.log(e.message);
    };

    map.on('locationfound', onLocationFound);
    map.on('locationerror', onLocationError);
    map.locate({ watch: true, enableHighAccuracy: true });
}

function onOverlayAdd(e) {
    this.append(e);
}

function onOverlayRemove(e) {
    this.detach(e);
}

// TODO(amwolff): allow only unique commits
class UserHistory {
    constructor() {
        this._params = new URL(window.location.href).searchParams;

        this._current_state = [];

        const state = history.state;
        if (state !== null) {
            this._params.delete('r');
            this._current_state = state;
            return;
        }

        if (this._params.has('r')) {
            const rts = this._getAvailableRoutesArray();
            this._params.get('r').split(',').forEach(p => {
                if (rts.includes(p) && !this._current_state.includes(p)) {
                    this._current_state.push(p);
                }
            });
            this._params.delete('r');
        }

        this._commitReplace();
    };

    _getAvailableRoutesArray() {
        const availableRoutesArray = [];
        availableRoutes.forEach(r => {
            availableRoutesArray.push(r.route);
        });
        return availableRoutesArray;
    };

    _getParamString() {
        const path = L.Util.getParamString({ r: this._current_state.sort() });

        const params = this._params.toString();
        if (params.length === 0) {
            return path;
        }
        return path.concat('&').concat(params);
    };

    _commitReplace() {
        history.replaceState(this._current_state, '', this._getParamString());
    };

    _commit() {
        history.pushState(this._current_state, '', this._getParamString());
    };

    _remove(element) {
        const idx = this._current_state.indexOf(element);
        if (idx !== -1) {
            this._current_state.splice(idx, 1);
        }
    };

    append(group) {
        if (group.name === '*') {
            this._current_state = this._getAvailableRoutesArray();
            this._commit();
            return;
        }

        this._current_state.push(group.name);
        this._commit();
    };

    detach(group) {
        if (group.name === '*') {
            this._current_state = [];
            this._commit();
            return;
        }

        this._remove(group.name);
        this._commit();
    };

    maybeAddGroups(map) {
        this._current_state.forEach(r => {
            vehiclesLayerGroups[r].addTo(map);
        });
    };

    onPop(map, event) {
        this._current_state = event.state;

        // Silence add/remove events so that they won't fire the propagation
        // chain where the history may get rewritten ("hazardous" situation).
        map.off('overlayadd', onOverlayAdd, this);
        map.off('overlayremove', onOverlayRemove, this);

        this._getAvailableRoutesArray().forEach(r => {
            if (this._current_state.includes(r)) {
                return;
            }
            vehiclesLayerGroups[r].removeFrom(map);
        });

        this._current_state.forEach(r => {
            vehiclesLayerGroups[r].addTo(map);
        });

        map.on('overlayadd', onOverlayAdd, this);
        map.on('overlayremove', onOverlayRemove, this);

        // TODO(amwolff): the dummy layer ('*') should also get cleaned up
    };
}

function initializeOverlays(map, userHistory) {
    availableRoutes.forEach(r => {
        vehiclesLayerGroups[r.route] = new L.LayerGroup();
    });
    map.on('overlayadd', onOverlayAdd, userHistory);
    map.on('overlayremove', onOverlayRemove, userHistory);
}

function addDummyLayerGroup(map, ctx) {
    // '*' is a dummy overlay used to enable all other overlays.
    // More info: https://github.com/Leaflet/Leaflet/issues/6400
    vehiclesLayerGroups['*'] = new L.LayerGroup();
    vehiclesLayerGroups['*'].on('add', () => {
        setTimeout(() => {
            map.off('overlayadd', onOverlayAdd, ctx);
            availableRoutes.forEach(r => {
                vehiclesLayerGroups[r.route].addTo(map);
            });
            map.on('overlayadd', onOverlayAdd, ctx);
        }, 0);
    });
    vehiclesLayerGroups['*'].on('remove', () => {
        setTimeout(() => {
            map.off('overlayremove', onOverlayRemove, ctx);
            availableRoutes.forEach(r => {
                vehiclesLayerGroups[r.route].removeFrom(map);
            });
            map.on('overlayremove', onOverlayRemove, ctx);
        }, 0);
    });
}

const endpointRoutes = 'http://localhost:8080/Routes';

function init() {
    const map = L.map('map', { attributionControl: false, center: [53.773056, 20.476111], zoom: 14 });

    L.tileLayer.wms('http://msipmo.olsztyn.eu/arcgis/services/msipmo_Plan/MapServer/WMSServer?', {
        tileSize: 256,
        minZoom: 9,
        maxZoom: 18,
        layers: '0,1,2,3,4,5,6,7,8,9,10,11,12',
        format: 'image/png',
    }).addTo(map);

    // L.control.attribution({
    //     prefix: false,
    //     position: 'bottomright',
    // }).addAttribution(
    //     '<a href="mailto:kontakt@autobusy.olsztyn.pl">Zgłoś błąd</a>' +
    //     ' / © <a href="https://www.mapbox.com/about/maps/">Mapbox</a>' +
    //     ' © <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>' +
    //     ' <strong><a href="https://www.mapbox.com/map-feedback/" target="_blank">Improve this map</a></strong>').addTo(map);

    setLocationTracking(map);

    fetch(endpointRoutes)
        .then(response => {
            return response.json();
        })
        .then(responseJSON => {
            availableRoutes = JSON.parse(JSON.stringify(responseJSON));

            const userHistory = new UserHistory();

            initializeOverlays(map, userHistory);
            addDummyLayerGroup(map, userHistory);

            userHistory.maybeAddGroups(map);

            window.onpopstate = L.bind(userHistory.onPop, userHistory, map);

            L.control.layers(null, vehiclesLayerGroups).addTo(map);

            refresh();
        });
}

window.onload = init;
