<?php
//Setup our REST framework library
include_once './epiphany/src/Epi.php';
include_once './http_response_code.php';
include_once './server/app.php';
include_once './server/user.php';
include_once './server/objective.php';
include_once './server/action.php';
include_once './server/reward.php';
include_once './server/type.php';
include_once './database.php';
Epi::setPath('base','./epiphany/src');
Epi::init('route');
//Epi::setSetting('exceptions',true);

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

/*****************************************************************
 * These functions are called to return http response codes      *
 * Use them to handle errors in your code cleanly                *
 *****************************************************************/

//Return 202 and optionally reply with a response string
function resourceCreated($response) {
	http_response_code(201);
	header('Content-Type: application/json');
	echo $response;
}

function resourceNotFound() {
	http_response_code(404);
	die("not found"); //TODO remove string. Only for debugging purposes.
}

function internalServerError($msg) {
	http_response_code(500);
	die($msg);
}

function malformedRequest() {
	http_response_code(400);
	die("malformed"); //TODO remove string. Only for debugging.
}
?>
