<?php
include_once('./objects/app.php');
class AppAPI {

    public static function CreateEndpoint() {
        $app = AppObject::buildAppFromObject(json_decode(file_get_contents('php://input')));
        AppAPI::create($app);
    }

    public static function MapUserEndpoint() {
        $uri = explode("/",$_SERVER['REQUEST_URI']);
        $app = $uri[2];
        $user = $uri[4];
        $map = MapObject::newMap($app,$user);
        $db = new Database();
        if(!$db->connect() || !$db->mapUserApp($map)) {
            internalServerError($db->getError());
        }
        resourceCreated("");
    }

    public static function MapObjectiveEndpoint() {
        $uri = explode("/",$_SERVER['REQUEST_URI']);
        $app = $uri[2];
        $objective = $uri[4];
        $map = MapObject::newMap($app,$objective);
        $db = new Database();
        if(!$db->connect() || !$db->mapObjectiveApp($map)) {
            internalServerError($db->getError());
        }
        resourceCreated("");
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
