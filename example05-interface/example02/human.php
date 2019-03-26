<?php

interface HumanInterface {
  public function run();
  public function say();
}

class Human implements HumanInterface {
  public $age;
  public $name;
  public function run() {
      echo "{$this->name}, I can run\n";
  }
  public function say() {
      echo "{$this->name}, I am {$this->age} years old\n";
  }
}

class Singer extends Human {
  public $collection;
  public function __construct($name, $age) {
    $this->name = $name;
    $this->age = $age;
  }
  public function sing() {
      echo "I can sing\n";
  }
}

class Student extends Human {
  public $lesson;
  public function __construct($name, $age) {
    $this->name = $name;
    $this->age = $age;
  }
  public function learn() {
      echo "I need learn {$lesson}\n";
  }
}

$foo = new Singer("foo", 31);
$foo->sing();
$foo->say();
$foo->run();

$bar = new Singer("bar", 32);
$bar->sing();
$bar->say();
$bar->run();
