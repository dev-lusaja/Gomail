#GOMAIL

Service that will allow us to connect to the MailGun API to send mails in a practical and fast way.

	* Multiple sending of emails.
	* Multiple answers.


Resuqest format
---------------

The requests are sent through the POST method, to the endPoint:

  >http://<ip>:<port>/api/v1/gomail

Every request must have the following structure:

~~~json
{
	"data" : [
		{
			"from" : "foo@barr.com",
			"subject" : "Test",
			"body" : "Hello World",
			"to" : "bar@foo.com"
		},
		{
			"from" : "foo2@bar2.com",
			"subject" : "Test 2",
			"body" : "Hello World 2",
			"to" : "bar2@foo2.com"
		}
	]
}
~~~

Response format
---------------

Once completed the shipments, GOMAIL will give a response with the following structure:

~~~json
{
  "responses": [
    {
      "id": "<20150817005339.xxxx.zzzz@mailgun.org>",
      "to": "foo@bar.com",
      "msg": "Queued. Thank you."
    },
    {
      "id": "<20150817005340.xxxx.zzzz@mailgun.org>",
      "to": "bar@foo.com",
      "msg": "Queued. Thank you."
    }
  ]
}
~~~



