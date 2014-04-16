<?php
include_once('./objects/user.php');
class UserAPI {
    public static function CreateEndpoint() {
        $app = UserObject::buildUserFromObject(json_decode(file_get_contents('php://input')));
	UserAPI::create($app);
    }

    public static function create($app) {
	$db = new Database();
	if(!$db->connect() || !$db->createUser($app)) {
		internalServerError($db->getError());
	}
	resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
    }
}
?>
