function display(arg) {
    console.log("clicked", arg)
    oldValue = document.getElementById('result').value;
    document.getElementById('result').value = oldValue + arg
};

function reset() {
    document.getElementById('result').value = ''
}

function del() {
    let oldValue = document.getElementById('result').value;
    document.getElementById('result').value = oldValue.slice(0, -1)
}

function calc() {
    let  uri = 'http://localhost:8088/calc/' + document.getElementById('result').value ;
    console.log(uri)
    
    fetch(uri, {
        method: 'GET',
        headers: {
            'Content-type': 'application/json'
        }
    })
    .then(function(response) {
        if (!response.ok) {
            console.log(error)
            document.getElementById('result').value = 'Unknown'
        }

       response.json().then(function (dataJson) {
           console.log(dataJson)
            document.getElementById('result').value = dataJson.value
       })
    })
    .catch(function(error) {
        console.log(error)
        document.getElementById('result').value = 'Unknown'
    })
}
