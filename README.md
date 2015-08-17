#GOMAIL

Es un microservicio que permitirá conectarnos al API de MailGun para realizar envíos de correos de una manera práctica y rápida, GOMAIL nos permite hacer:

	* Múltiple solicitud de envíos.
	* Envío de correos múltiples.
	* Respuesta a la(s) solicitudes.
	


Formato de la solicitud
-----------------------
Las solicitudes se envían a través del metodo POST, hacía el endPoint:

  >http://localhost:5000/api/v1/Gomail

Toda solicitud debe tener la siguiente estructura:

~~~json
{
	"data" : [
		{
			"from" : "yyyyyyy@hotmail.com",
			"subject" : "Test",
			"body" : "Hello World",
			"to" : "xxxxxxxxx@gmail.com"
		},
		{
			"from" : "yyyyyyy@hotmail.com",
			"subject" : "Test2",
			"body" : "Hello World2",
			"to" : "zzzzzzzz@gmail.com"
		}
	]
}
~~~

Formato de la respuesta 
-----------------------
Una vez culminado el o los envíos, GOMAIL dara una respuesta con la siguiente estructura:

~~~json
{
  "responses": [
    {
      "id": "<20150817005339.6788.94173@mailgun.org>",
      "to": "xxxxxxxx@gmail.com",
      "msg": "Queued. Thank you."
    },
    {
      "id": "<20150817005340.47422.6102@mailgun.org>",
      "to": "zzzzzzzz@gmail.com",
      "msg": "Queued. Thank you."
    }
  ]
}
~~~



