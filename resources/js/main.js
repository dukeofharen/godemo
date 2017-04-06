var broken = "img/egg-broken.svg";
var maxId;

document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("egg1").style.marginTop = rnd(5) + "em";
    document.getElementById("egg2").style.marginTop = rnd(15) + "em";
    document.getElementById("egg3").style.marginTop = rnd(5) + "em";
    document.getElementById("egg4").style.marginTop = rnd(15) + "em";
    document.getElementById("egg5").style.marginTop = rnd(5) + "em";
});

function rnd(range) {
    return Math.floor((Math.random() * range));
}

function egg1() {
    httpGet("/api/greet/", function(data) {
        document.getElementById("title").textContent = data;
        document.getElementById("egg1").src = "img/egg1.svg";
    }, function () {
        document.getElementById("egg1").src = broken;
    });
}

function egg2() {
    getJson("GET", "/api/greet/", null, function(data) {
        if (data.title && data.subtitle) {
            document.getElementById("title").textContent = data.title;
            document.getElementById("subtitle").textContent = data.subtitle;
            document.getElementById("egg2").src = "img/egg2.svg";
        } else {
            document.getElementById("egg2").src = broken;
        }
    }, function () {
        document.getElementById("egg2").src = broken;
    });
}

function egg3() {
    var nums = [];
    var count = 50 + rnd(50);
    var start = rnd(50);
    var end = start + rnd(count - start);
    var answer = 0;
    for (var i = 0; i < count; i++) {
        var r = rnd(100);
        nums.push(r);
        if (i >= start && i < end) {
            answer += r
        }
    }
    var requestBody = {start: start, end: end, numbers: nums};
    getJson("POST", "/api/sum/", requestBody, function(data) {
        if (data.answer === answer && data.contributers === (end - start)) {
            document.getElementById("egg3").src = "img/egg3.svg";
        } else {
            document.getElementById("egg3").src = broken;
        }
    }, function () {
        document.getElementById("egg3").src = broken;
    });
}

function egg4() {
    var name = document.getElementById("name").value;
    var message = document.getElementById("message").value;
    if (!name || !message) {
        document.getElementById("egg4").src = broken;
    }
    var requestBody = {
        name: name,
        message: message
    };
    getJson("POST", "/api/store/", requestBody, function(data) {
        if (data.id) {
            maxId = data.id;
            document.getElementById("egg4").src = "img/egg4.svg";
        } else {
            document.getElementById("egg4").src = broken;
        }
    }, function () {
        document.getElementById("egg4").src = broken;
    });
}

function egg5() {
    if (!maxId) {
        document.getElementById("egg5").src = broken;
        return;
    }

    var id = 1 + rnd(maxId);
    getJson("GET", "/api/store/"+ id, null, function(data) {
        if (data.name && data.message) {
            document.getElementById("subtitle").textContent = data.message + " -- " + data.name;
            document.getElementById("egg5").src = "img/egg5.svg";
        } else {
            document.getElementById("egg5").src = broken;
        }
    }, function() {
        document.getElementById("egg5").src = broken;
    });
}

function httpGet(url, callbackOK, callbackError) {
    var req = new XMLHttpRequest();
    req.onreadystatechange = function() {
        if (req.readyState === XMLHttpRequest.DONE && req.status === 200) {
            callbackOK(req.responseText);
        } else {
            callbackError();
        }
    };
    req.open("GET", url, true);
    req.setRequestHeader("Accept", "text/plain");
    req.send(null);
}

function getJson(action, url, data, callbackOK, callbackError) {
    var req = new XMLHttpRequest();
    req.onreadystatechange = function() {
        if (req.readyState === XMLHttpRequest.DONE && req.status === 200
                && req.getResponseHeader("Content-Type").indexOf("application/json") !== -1) {
            callbackOK(JSON.parse(req.responseText));
        } else {
            callbackError();
        }
    };
    req.open(action, url, true);
    req.setRequestHeader("Accept", "application/json");
    req.send(JSON.stringify(data));
}
