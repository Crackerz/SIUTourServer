<?php
class MapObject {
    //For app.mapUser it would be
    //firstId=appId
    //seconId=userId
    public $mapId;
    public $firstId;
    public $secondId;

    public static function newMap($first,$second,$id=null) {
        $result = new MapObject();
        $result->firstId = $first;
        $result->secondId = $second;
        if(!is_null($id))
            $result->mapId = $id;
        return $result;
    }
}
?>
