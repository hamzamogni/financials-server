<?php

namespace App\Services;

use App\Repositories\Interfaces\UserRepositoryInterface;

class AuthService
{
    protected $userRepository;

    public function __construct(UserRepositoryInterface $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    /**
     * @param array $data
     * @return array
     */
    public function signUp(array $data): array
    {
        $user = $this->userRepository->create($data);

        //$this->userRepository->assignRole($user, $data["role"]);

        $token = $this->userRepository->createToken($user);

        return [
            "user" => $user,
            "token" => $token
        ];
    }
}
