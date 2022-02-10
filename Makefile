
get		:
		curl -L -X GET 'localhost:8080/' -H 'Content-Type: application/json' \
 		--data-raw '{"url": "1a3_09Ezhi"}'

post	:
		curl -L -X POST 'localhost:8080/' -H 'Content-Type: application/json' \
 		--data-raw '{"url": "https://www.alexedwards.net/blog/working-with-rediss"}'