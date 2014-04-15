<?php
class App {
    public $Id;
    public $Name;
    public $Token;
    public $Objetives;
    public $Users;

    public static function create() {
        $app = App::buildAppFromObject(json_decode(file_get_contents('php://input')));
        echo '{"Result":{"Id":10}}';
    }

    public static function buildAppFromObject(&$obj) {
        $result = new App();
        //Grab all property names of this object and populate them if
        //the object had the same names declared
        $properties = array_keys(get_object_vars($result));
        foreach($properties as $property) {
            if(property_exists($obj,$property)) {
                $result->$property = $obj->$property;
            }
        }
        return $result;
    }
}
?>
