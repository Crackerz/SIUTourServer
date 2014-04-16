<?php
class UserObject {
    public $Id;
    public $Fname;
    public $Lname;
    public $Birthday;
    public $UserId;
    public $Apps;
    public $Objectives;

    public static function buildUserFromObject($obj) {
        $result = new UserObject();
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
