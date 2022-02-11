
get		:
		curl -L -X GET 'localhost:8080/' -H 'Content-Type: application/json' \
 		--data-raw '{"url_short": "t28ZxoYX"}'

post	:
		curl -L -X POST 'localhost:8080/' -H 'Content-Type: application/json' \
 		--data-raw '{"url_long": "https://www.alexewearethe.loh/gg"}'