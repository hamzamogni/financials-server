<?php

namespace Tests\Unit;

use App\Models\User;
use Illuminate\Support\Facades\Hash;
use PHPUnit\Framework\TestCase;

class PasswordHashingTest extends TestCase
{
    /**
     * A basic unit test example.
     *
     * @return void
     */
    public function testPasswordGetsHashedWhenSet()
    {
        Hash::shouldReceive("make")->once()->andReturn("hashed");

        $user = new User();
        $user->password = "test";

        $this->assertEquals("hashed", $user->password);
    }
}
