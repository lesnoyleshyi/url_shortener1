
get		:
			curl -L -X GET 'localhost:8080/' -H 'Content-Type: application/json' \
			--data-raw '{"url_short": "1tACP_QkAE"}'

post	:
		curl -L -X POST 'localhost:8080/' -H 'Content-Type: application/json' \
 		--data-raw '{"url_long": "https://www.alexeweefwefarethe.loh/gg/ewfewfw"}'