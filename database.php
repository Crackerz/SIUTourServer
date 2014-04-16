<?php
require('./credentials.php');

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
		$query = "INSERT INTO `".$GLOBALS["APP_TABLE"]."` (`name`,`token`,`status`) VALUES ('".$app->Name."','".hash("sha512",rand())."',0)";
		return $this->connection->query($query);
	}
	
	public function createUser($user) {
		$query = "INSERT INTO `".$GLOBALS["USER_TABLE"]."` (`fname`,`lname`,`bday`,`userId`) VALUES ('".$user->Fname."','".$user->Lname."','".$user->Birthday."','".$user->UserId."')";
		return $this->connection->query($query);
	}
}
?>
