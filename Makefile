
get		:
		curl -L -X GET 'localhost:8080/' -H 'Content-Type: application/json' \
			--data-raw '{"url_short": "GkNk7Plccu"}'

post	:
		curl -L -X POST 'localhost:8080/' -H 'Content-Type: application/json' \
 		--data-raw '{"url_long": "https://www.github.com/lesnoyleshyi"}'

get_oracle	:
		curl -L -X GET 'http://132.226.200.167:8080/' -H 'Content-Type: application/json' \
        			--data-raw '{"url_short": "4UloxNyc4J"}'

post_oracle	:
		curl -L -X POST 'http://132.226.200.167:8080/' -H 'Content-Type: application/json' \
         		--data-raw '{"url_long": "https://www.github.com/lesnoyleshyi"}'