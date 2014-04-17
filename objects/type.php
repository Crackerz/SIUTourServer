<?php
class TypeObject {
    public $Id;
    public $Name;
    public $Description;

    public static function buildActionFromObject($obj) {
        $result = new TypeObject();
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

    public static function buildRewardFromObject($obj) {
        $result = new TypeObject();
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
