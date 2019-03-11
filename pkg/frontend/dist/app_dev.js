var iconTemplate='<svg class="outerVehicleIcon" width="26" height="26">\n<path transform="rotate({rotate} 13 19)" fill="{fill}" stroke="{stroke}" stroke-width="2"\nstroke-dasharray="{stroke_dasharray}" d="\nM26\n19c0-2.2-0.6-4.4-1.6-6.2C22.2\n8.8\n13\n0\n13\n0S3.8\n8.7\n1.6\n12.8c-1\n1.8-1.6\n4-1.6\n6.2c0\n7.2\n5.8\n13\n13\n13\nS26\n26.2\n26\n19 Z"/>\n<g>\n<text x="13" y="19" font-family="sans-serif" font-size="12px" fill="white"\ntext-anchor="middle" alignment-baseline="central">{route}\n</text>\n</g>\nSorry, your browser does not support inline SVG.\n</svg>',
popupTemplate="<b>Linia nr {route}</b><br>Numer kursu: {trip_id}<br>Kierunek: {description}<br>{state}: {variance}",InternalVehicle=function(a){(this.stall=0===a.route.length)?(this.azimuth=0,this.description=a.next_description,this.route=a.next_route,this.trip_id=a.next_trip_id):(this.azimuth=180+a.azimuth,this.description=a.description,this.route=a.route,this.trip_id=a.trip_id);this.latitude=a.latitude;this.longitude=a.longitude;this.variance=a.variance};
InternalVehicle.getIconBodyColor=function(a){switch(a){case "N01":case "N02":return"rgba(0, 0, 0, 0.9)";case "1":case "2":case "3":return"rgba(227, 30, 30, 0.9)"}return"rgba(0, 157, 210, 0.9)"};InternalVehicle.getIconBorderColor=function(a){return-120<a?"white":-600<a?"yellow":"orange"};InternalVehicle.getIconBorderStyle=function(a){return a?"2,2":"0,0"};
InternalVehicle.hrSeconds=function(a){if(60<a){var b=Math.floor(a/60);a-=60*b;if(60<b){var c=Math.floor(b/60);return L.Util.template("{HH}h {MM}m {SS}s",{HH:c,MM:b-60*c,SS:a})}return L.Util.template("{MM}m {SS}s",{MM:b,SS:a})}return L.Util.template("{SS}s",{SS:a})};
InternalVehicle.prototype.getIconHTML=function(){var a={rotate:this.azimuth,fill:InternalVehicle.getIconBodyColor(this.route),stroke:InternalVehicle.getIconBorderColor(this.variance),stroke_dasharray:InternalVehicle.getIconBorderStyle(this.stall),route:this.route};return L.Util.template(iconTemplate,a)};
InternalVehicle.prototype.getPopupContent=function(){var a={route:this.route,trip_id:this.trip_id,description:this.description};this.stall?0<this.variance?(a.state="Czas do odjazdu",a.variance=InternalVehicle.hrSeconds(this.variance)):(a.state="Op\u00f3\u017anienie odjazdu",a.variance=InternalVehicle.hrSeconds(-1*this.variance)):0<this.variance?(a.state="Przed czasem",a.variance=InternalVehicle.hrSeconds(this.variance)):(a.state="Op\u00f3\u017anienie",a.variance=InternalVehicle.hrSeconds(-1*this.variance));
return L.Util.template(popupTemplate,a)};InternalVehicle.prototype.getLeafletDivIcon=function(){return L.divIcon({html:this.getIconHTML(),className:"innerVehicleIcon"})};InternalVehicle.prototype.getLeafletLatLng=function(){return L.latLng(this.latitude,this.longitude)};InternalVehicle.prototype.getLeafletMarker=function(){var a={icon:this.getLeafletDivIcon(),title:this.route,alt:this.trip_id,riseOnHover:!0};return L.marker(this.getLeafletLatLng(),a).bindPopup(this.getPopupContent(),{autoPan:!1})};
function serializeVehicles(a){var b=[];a.forEach(function(a){b.push(new InternalVehicle(a))});return b}function updateMarker(a,b){a.setIcon(b.getLeafletDivIcon());a.setLatLng(b.getLeafletLatLng());var c=a.getPopup();c._content=b.getPopupContent();c.update()}var availableRoutes,vehiclesLayerGroups=[];
function insertOnMap(a){var b=serializeVehicles(a);availableRoutes.forEach(function(a){vehiclesLayerGroups[a.route].eachLayer(function(c){var d=!1;b.forEach(function(a){a.trip_id===c.options.alt&&(d=!0)});d||vehiclesLayerGroups[a.route].removeLayer(c)});b.forEach(function(b){if(b.route===a.route){var c=!1;vehiclesLayerGroups[a.route].eachLayer(function(a){a.options.alt===b.trip_id&&(c=!0,updateMarker(a,b))});c||vehiclesLayerGroups[a.route].addLayer(b.getLeafletMarker())}})})}
function planNextRefresh(a){a=7E3-(Date.now()-a);0>a?-6E3<a?setTimeout(refresh,1E3):console.log("Refreshing the map has been stalled. This can happen when there's a problem with the data source or with your clock. Reload the page to start the refreshing again."):setTimeout(refresh,a)}var endpointVehicles="http://localhost:8080/Vehicles";
function refresh(){fetch(endpointVehicles).then(function(a){var b=new Date(a.headers.get("Last-Modified"));planNextRefresh(b);return a.json()}).then(function(a){insertOnMap(JSON.parse(JSON.stringify(a)))})}
function setLocationTracking(a){var b;a.on("locationfound",function(c){var d=c.accuracy/2;a.hasLayer(b)?b.setLatLng(c.latlng).setRadius(d):(b=L.circle(c.latlng,{radius:d,color:"#FF6C00"}).addTo(a),a.flyToBounds(b.getBounds(),{maxZoom:17}))});a.on("locationerror",function(a){console.log(a.message)});a.locate({watch:!0,enableHighAccuracy:!0})}function onOverlayAdd(a){this.append(a)}function onOverlayRemove(a){this.detach(a)}
var UserHistory=function(){var a=this;this._params=(new URL(window.location.href)).searchParams;this._current_state=[];var b=history.state;if(null!==b)this._params["delete"]("r"),this._current_state=b;else{if(this._params.has("r")){var c=this._getAvailableRoutesArray();this._params.get("r").split(",").forEach(function(b){c.includes(b)&&!a._current_state.includes(b)&&a._current_state.push(b)});this._params["delete"]("r")}this._commitReplace()}};
UserHistory.prototype._getAvailableRoutesArray=function(){var a=[];availableRoutes.forEach(function(b){a.push(b.route)});return a};UserHistory.prototype._getParamString=function(){var a=L.Util.getParamString({r:this._current_state.sort()}),b=this._params.toString();return 0===b.length?a:a.concat("&").concat(b)};UserHistory.prototype._commitReplace=function(){history.replaceState(this._current_state,"",this._getParamString())};
UserHistory.prototype._commit=function(){history.pushState(this._current_state,"",this._getParamString())};UserHistory.prototype._remove=function(a){a=this._current_state.indexOf(a);-1!==a&&this._current_state.splice(a,1)};UserHistory.prototype.append=function(a){"*"===a.name?this._current_state=this._getAvailableRoutesArray():this._current_state.push(a.name);this._commit()};UserHistory.prototype.detach=function(a){"*"===a.name?this._current_state=[]:this._remove(a.name);this._commit()};
UserHistory.prototype.maybeAddGroups=function(a){this._current_state.forEach(function(b){vehiclesLayerGroups[b].addTo(a)})};
UserHistory.prototype.onPop=function(a,b){var c=this;this._current_state=b.state;a.off("overlayadd",onOverlayAdd,this);a.off("overlayremove",onOverlayRemove,this);this._getAvailableRoutesArray().forEach(function(b){c._current_state.includes(b)||vehiclesLayerGroups[b].removeFrom(a)});this._current_state.forEach(function(b){vehiclesLayerGroups[b].addTo(a)});a.on("overlayadd",onOverlayAdd,this);a.on("overlayremove",onOverlayRemove,this)};
function initializeOverlays(a,b){availableRoutes.forEach(function(a){vehiclesLayerGroups[a.route]=new L.LayerGroup});a.on("overlayadd",onOverlayAdd,b);a.on("overlayremove",onOverlayRemove,b)}
function addDummyLayerGroup(a,b){vehiclesLayerGroups["*"]=new L.LayerGroup;vehiclesLayerGroups["*"].on("add",function(){setTimeout(function(){a.off("overlayadd",onOverlayAdd,b);availableRoutes.forEach(function(b){vehiclesLayerGroups[b.route].addTo(a)});a.on("overlayadd",onOverlayAdd,b)},0)});vehiclesLayerGroups["*"].on("remove",function(){setTimeout(function(){a.off("overlayremove",onOverlayRemove,b);availableRoutes.forEach(function(b){vehiclesLayerGroups[b.route].removeFrom(a)});a.on("overlayremove",
onOverlayRemove,b)},0)})}var endpointRoutes="http://localhost:8080/Routes";
function init(){var a=L.map("map",{attributionControl:!1,center:[53.773056,20.476111],zoom:14});L.tileLayer.wms("http://msipmo.olsztyn.eu/arcgis/services/msipmo_Plan/MapServer/WMSServer?",{tileSize:256,minZoom:9,maxZoom:18,layers:"0,1,2,3,4,5,6,7,8,9,10,11,12",format:"image/png"}).addTo(a);setLocationTracking(a);fetch(endpointRoutes).then(function(a){return a.json()}).then(function(b){availableRoutes=JSON.parse(JSON.stringify(b));b=new UserHistory;initializeOverlays(a,b);addDummyLayerGroup(a,b);b.maybeAddGroups(a);
window.onpopstate=L.bind(b.onPop,b,a);L.control.layers(null,vehiclesLayerGroups).addTo(a);refresh()})}window.onload=init;
