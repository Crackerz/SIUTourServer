<?php
class RewardObject {
    public $Id;
    public $Name;
    public $Type;
    public $Value;

    public static function buildRewardFromObject($obj) {
        $result = new RewardObject();
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
