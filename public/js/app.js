$(document).ready(function () {
    let loc = window.location;
    let uri = 'ws:';

    if (loc.protocol === 'https:') {
        uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += loc.pathname + 'ws';

    let ws = new WebSocket(uri);

    ws.onopen = function () {
        console.log('Connected')
    };

    ws.onmessage = function (evt) {
        let obj = JSON.parse(evt.data);
        let message = `<li class="list-group-item"><div class="row"><div class="col-sm">"${ obj.body.message }"</div></div></li>`;
        $("#output").append(message);
    };
});