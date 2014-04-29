<?php
include_once('./objects/objective.php');
class ObjectiveAPI {
	public static function CreateEndpoint() {
		$objective = ObjectiveObject::buildObjectiveFromObject(json_decode(file_get_contents('php://input')));
		ObjectiveAPI::create($objective);
	}

	public static function create($objective) {
		$db = new Database();
		if(!$db->connect() || !$db->createObjective($objective)) {
			internalServerError($db->getError());
		}

		resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
	}

    public static function MapActionEndpoint() {
        $uri = explode("/",$_SERVER['REQUEST_URI']);
        $objective = $uri[2];
        $action = $uri[4];
        $map = MapObject::newMap($objective,$action);
        $db = new Database();
        if(!$db->connect() || !$db->mapActionObjective($map)) {
            internalServerError($db->getError());
        }
        resourceCreated("");
    }

    public static function MapRewardEndpoint() {
        $uri = explode("/",$_SERVER['REQUEST_URI']);
        $objective = $uri[2];
        $reward = $uri[4];
        $map = MapObject::newMap($objective,$reward);
        $db = new Database();
        if(!$db->connect() || !$db->mapRewardObjective($map)) {
            internalServerError($db->getError());
        }
        resourceCreated("");
    }

}
