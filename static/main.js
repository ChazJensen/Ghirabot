function getCard(cardName) {
	var reqUrl = "http://localhost:8080/lookup/" + cardName.replace(" ", "+")
	var responseText = "";

	var xmlHttp = new XMLHttpRequest();

	xmlHttp.onreadystatechange = function() {
		if (this.readyState == 4) {
			console.log(this.responseText)
			responseText = this.responseText;
		}
	};
	
	xmlHttp.open("GET", reqUrl, false);

	xmlHttp.send();

	return xmlHttp.responseText;
}
