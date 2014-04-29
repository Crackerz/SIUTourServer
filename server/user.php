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

    public static function MapObjectiveEndpoint() {
        $uri = explode("/",$_SERVER['REQUEST_URI']);
        $user = $uri[2];
        $objective = $uri[4];
        UserAPI::mapObjective($user,$objective);
        resourceCreated("");
    }

    public static function mapObjective($userId,$objectiveId) {
        $map = MapObject::newMap($userId,$objectiveId);
        $db = new Database();
        if(!$db->connect() || !$db->mapObjectiveUser($map)) {
            internalServerError($db->getError());
        }
    }


}
?>
