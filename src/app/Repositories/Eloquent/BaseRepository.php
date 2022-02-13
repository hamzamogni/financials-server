<?php

namespace App\Repositories\Eloquent;

use App\Repositories\Interfaces\EloquentRepositoryInterface;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Model;

class BaseRepository implements EloquentRepositoryInterface
{
    protected $model;

    public function __construct(Model $model)
    {
        $this->model = $model;
    }

    public function all(): Collection
    {
        return $this->model::all();
    }

    public function create(array $attributes): Model
    {
        return $this->model->create($attributes);
    }

    /**
     * @param              $id
     * @codeCoverageIgnore
     */
    public function find($id): Model
    {
        return $this->model->findOrFail($id);
    }

    public function update(Model $model, array $attributes): Model
    {
        $model->update($attributes);
        return $model->fresh();
    }

    public function destroy(Model $model): bool
    {
        $model->delete();
        return true;
    }
}
