<?php

namespace Tests\Feature;

use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Foundation\Testing\WithFaker;
use Tests\TestCase;

class AuthControllerTest extends TestCase
{
    use RefreshDatabase, WithFaker;

    public function testUserCanSignup()
    {
        $userData = [
            "name" => $this->faker->name,
            "email" => $this->faker->safeEmail,
            "password" => $this->faker->password,
        ];

        $response = $this->postJson("/api/sign-up", $userData);

        $this->assertDatabaseHas(
            "users",
            [
            "name" => $userData["name"],
            "email" => $userData["email"],
            ]
        );

        $response->assertCreated();

        $response->assertJsonStructure(["user", "token"]);
    }
}
