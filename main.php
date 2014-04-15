<?php

//Setup our REST framework library
include_once './epiphany/src/Epi.php';
include_once './http_response_code.php';
Epi::setPath('base','./epiphany/src');
Epi::init('route');
Epi::setSetting('exceptions',true);

//Load in our API configuration file
Epi::setPath('config','./');
getRoute()->load('api.ini');
getRoute()->run();

class test {
	public static function runme() {
		http_response_code(200);
		die("ERROR");
	}
}

?>
