<?php

namespace App\Repositories\Interfaces;

use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Model;

interface EloquentRepositoryInterface
{
    public function all(): Collection;

    public function create(array $attributes): Model;

    public function find($id): ?Model;

    public function update(Model $model, array $attributes): Model;

    public function destroy(Model $model): bool;
}
