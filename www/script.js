'use strict';


function convertToJson() {
let value = document.getElementById("textareaID").value;
convertToJsonGo(value)
}


function setValue(value){
    console.log(value)
    alert(value)
    document.getElementById('textareaResult').innerHTML = value;
}
