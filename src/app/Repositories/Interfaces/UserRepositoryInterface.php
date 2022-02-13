<?php


namespace App\Repositories\Interfaces;

use App\Models\User;

interface UserRepositoryInterface
{
    public function getByEmail(string $email): ?User;

    // public function assignRole(User $user, string $role);

    public function createToken(User $user);
}
