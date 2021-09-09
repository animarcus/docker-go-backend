async function post() {
    const url = 'http://127.0.0.1:8080/post/';
    const data = {
        "title": document.getElementById("postTitle").value,
        "content": document.getElementById("postContent").value
    }
    const params = {
        body: JSON.stringify(data),
        method: "POST"
    }

    fetch(url, params)
    .then(data => { 
        console.log(data);
        // return data.json()
    })
    .then(res => {
        console.log(res) })
}


async function get() {
    let index = document.getElementById("getId").value;
    if (index === "") index = 1;
    console.log(index);
    const url = 'http://127.0.0.1:8080/post/' + index;
    const params = {
        // headers: {
        //     "content-type": "application/json; charset=UTF-8"
        // },
        // body: data,
        method: "GET"
    }

    fetch(url, params)
    .then(data => { 
        console.log(data);
        return data.json() })
    .then(res => {
        console.log(res) })
}

async function put() {
    const inputId = document.getElementById("putId").value;
    const url = 'http://127.0.0.1:8080/post/' + inputId;
    const data = {
        "title": document.getElementById("putTitle").value,
        "content": document.getElementById("putContent").value
    }
    console.log(inputId, data, url)
    const params = {
        // headers: {
        //     "content-type": "application/json; charset=UTF-8"
        // },
        mode: "cors",
        body: JSON.stringify(data),
        method: "PUT"
    }
    // console.log(params)

    fetch(url, params)
    .then(data => { 
        console.log(data);
        // return data.json()
    })
    .then(res => {
        console.log(res) })
}


async function remove() {
    let index = document.getElementById("deleteId").value;
    if (index === "") index = 1;
    // console.log(index);
    const url = 'http://127.0.0.1:8080/post/' + index;
    const params = {
        // headers: {
        //     "content-type": "application/json; charset=UTF-8"
        // },
        // body: data,
        method: "DELETE"
    }

    fetch(url, params)
    .then(data => { 
        console.log(data);
        return data.json() })
    .then(res => {
        console.log(res) })
}


