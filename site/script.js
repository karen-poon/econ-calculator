var siteName = "https://ece472-calculator.netlify.app/.netlify/functions/compute?"

// getQueryString is a function that generates the url query string
function getQueryString() {
    var selection = document.getElementById("selection");
    // record the active selected equation
    var selected = selection.getElementsByClassName("selectBtn active")[0].id;
    // record their i, j(if exist), n values
    let itext = document.getElementById('iText').value;
    let ntext = document.getElementById('nText').value;
    let jtext = "";
    if (document.getElementById('jText') !== null) {
        jtext = document.getElementById('jText').value;
    }
    // store in a data struct
    let data = { type: selected, i: itext, j: jtext, n: ntext };
    // encode query string in url
    let queryString = Object.keys(data).map((key) => {
        return encodeURIComponent(key) + '=' + encodeURIComponent(data[key])
    }).join('&');

    return queryString;
}

function getDisplayResult(responseText) {
    let responseBody = JSON.parse(responseText);

    let flag = 0;
    let displayResult = "Invalid input for ";

    if (responseBody.iStatus.localeCompare("error") == 0) {
        flag += 1;
    }
    if (responseBody.jStatus.localeCompare("error") == 0) {
        flag += 10;
    }
    if (responseBody.nStatus.localeCompare("error") == 0) {
        flag += 100;
    }

    if (flag == 1) {                // 001
        displayResult += "i";
    } else if (flag == 10) {        // 010
        displayResult += "j";
    } else if (flag == 11) {        // 011
        displayResult += "i, j";
    } else if (flag == 100) {       // 100
        displayResult += "n";
    } else if (flag == 101) {       // 101
        displayResult += "i, n";
    } else if (flag == 110) {       // 110
        displayResult += "j, n";
    } else if (flag == 111) {       // 111
        displayResult += "i, j, n";
    } else { // flag = 0
        displayResult = responseBody.result;
    }
    return displayResult;
}

// send request to server when we hit the submit button
let submitBtn = document.getElementById('submit');
submitBtn.addEventListener ("click", function() {
    
    let queryString = getQueryString();
    // console.log(queryString);

    // send a GET request
    var submitRequest = new XMLHttpRequest();
    // var serverResponse;
    submitRequest.open("GET", siteName+queryString, true);
    submitRequest.setRequestHeader("Content-Type", "application/json");
    submitRequest.setRequestHeader("Access-Control-Allow-Origin", "*");
    submitRequest.send();
    submitRequest.onreadystatechange = function() {
        if (submitRequest.readyState == 4 && submitRequest.status == 200) { 
            var displayResult = getDisplayResult(submitRequest.responseText);

            document.getElementById('result').value = displayResult
        }
    };
});

// deals with the equation that we select
// selected button should remain active
// j input text box appears when we click (P|A i, j, n) or (F|A i, j, n)
var selection = document.getElementById("selection");
var selectBtns = selection.getElementsByClassName("selectBtn");
var jContainer = document.getElementById("j");
for (var i = 0; i < selectBtns.length; i++) {
    selectBtns[i].addEventListener("click", function() {
        jContainer.textContent = "";
        // if we don't add this, the jBox will be added many times
        let jBox = document.getElementById("jText");
        if (jBox) {           
            jContainer.removeChild(jBox);
        }

        var current = document.getElementsByClassName("active");
        //remove active from the previously active selectBtn
        current[0].className = current[0].className.replace(" active", "");
        // make the current one active
        this.className += " active";

        // add the jBox if it's (P|A i, j, n) or (F|A i, j, n)
        if (current[0].id == 9 || current[0].id == 10) {
            jContainer.textContent = "j: ";
            let jBox = document.createElement('input');
            jBox.setAttribute("id", "jText");
            jBox.setAttribute("placeholder", " e.g. 0.2");
            jContainer.appendChild(jBox);
        }
  });
}
