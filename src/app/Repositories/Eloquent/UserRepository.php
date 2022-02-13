<?php

namespace App\Repositories\Eloquent;

use App\Models\User;
use App\Repositories\Interfaces\UserRepositoryInterface;

class UserRepository extends BaseRepository implements UserRepositoryInterface
{
    public function __construct(User $model)
    {
        parent::__construct($model);
    }

    public function getByEmail(string $email): ?User
    {
        return User::where("email", $email)->first();
    }

    public function createToken(User $user)
    {
        return $user->createToken("user-string")->plainTextToken;
    }
}
