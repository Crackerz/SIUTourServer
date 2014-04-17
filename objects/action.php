<?php
class ActionObject {
	public $Id;
	public $Name;
	public $Type;
	public $Value;
    
	public static function buildActionFromObject($obj) {
        $result = new ActionObject();
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
