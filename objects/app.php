<?php
class AppObject {
    public $Id;
    public $Name;
    public $Token;
    public $Objetives;
    public $Users;

    public static function buildAppFromObject(&$obj) {
        $result = new AppObject();
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
