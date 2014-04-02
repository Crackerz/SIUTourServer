<?php

//Setup our REST framework library
include_once './epiphany/src/Epi.php';
Epi::setPath('base','./epiphany/src');
Epi::init('route');

//Load in our API configuration file
Epi::setPath('config','./');
getRoute()->load('api.ini');
getRoute()->run();

class test {
	public static function runme() {
		echo("Hello World!");
		http_response_code(500);
		die("ERROR");
	}
}

?>
