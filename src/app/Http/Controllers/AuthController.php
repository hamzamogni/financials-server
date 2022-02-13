<?php

namespace App\Http\Controllers;

use App\Services\AuthService;
use Illuminate\Http\Request;

class AuthController extends Controller
{
    protected $authService;

    public function __construct(AuthService $authService)
    {
        $this->authService = $authService;
    }

    public function signUp(Request $request)
    {
        $userWithToken = $this->authService->signUp($request->all());

        return response(
            $userWithToken,
            201
        );
    }
}
