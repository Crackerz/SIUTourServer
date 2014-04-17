<?php
require('./credentials.php');

//This php script will be included by another, all objects passed in are defined by the calling script.
//TODO: PREVENT INJECTION

class Database {
	public $connection;

	public function connect() {
		$this->connection = mysqli_connect($GLOBALS["SERVER_ADDRESS"],$GLOBALS["SERVER_USER"],$GLOBALS["SERVER_PASSWORD"],$GLOBALS["SERVER_DATABASE"]);
		if(mysqli_connect_errno($this->connection)) {
			$this->connection = null;
			return false;
		}
		return true;
	}

	public function getError() {
		return $this->connection->error;
	}

	public function createApp($app) {
		$query = "INSERT INTO `".$GLOBALS["APP_TABLE"]."` (`name`,`token`) VALUES ('".$app->Name."','".hash("sha512",rand())."')";
		return $this->connection->query($query);
	}
	
	public function createUser($user) {
		$query = "INSERT INTO `".$GLOBALS["USER_TABLE"]."` (`fname`,`lname`,`bday`,`userId`) VALUES ('".$user->Fname."','".$user->Lname."','".$user->Birthday."','".$user->UserId."')";
		return $this->connection->query($query);
	}

	public function createObjective($objective) {
		$query = "INSERT INTO `".$GLOBALS["OBJECTIVE_TABLE"]."` (`name`,`description`) VALUES ('".$objective->Name."','".$objective->Description."')";
		return $this->connection->query($query);
	}

	public function createAction($action) {
		$query = "INSERT INTO `".$GLOBALS["ACTION_TABLE"]."` (`name`,`type`,`value`) VALUES ('".$action->Name."',".$action->Type.",'".$action->Value."')";
		return $this->connection->query($query);
	}

	public function createActionType($actionType) {
		$query = "INSERT INTO `".$GLOBALS["ACTION_TYPE_TABLE"]."` (`name`, `description`) VALUES ('".$actionType->Name."','".$actionType->Description."')";
		return $this->connection->query($query);
	}

	public function createReward($reward) {
		$query = "INSERT INTO `".$GLOBALS["REWARD_TABLE"]."` (`name`,`type`,`value`) VALUES ('".$reward->Name."','".$reward->Type."','".$reward->Value."')";
		return $this->connection->query($query);
	}

	public function createRewardType($rewardType) {
		$query = "INSERT INTO `".$GLOBALS["REWARD_TYPE_TABLE"]."` (`name`,`description`) VALUES ('".$rewardType->Name."','".$rewardType->Description."')";
		return $this->connection->query($query);
	}
}
?>
