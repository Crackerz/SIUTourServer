<?php
include_once('./objects/app.php');
class AppAPI {
    public static function CreateEndpoint() {
        $app = AppObject::buildAppFromObject(json_decode(file_get_contents('php://input')));
	AppAPI::create($app);
    }

    public static function create($app) {
	$db = new Database();
	if(!$db->connect() || !$db->createApp($app)) {
		internalServerError($db->getError());
	}
	
	resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
    }
}
?>
