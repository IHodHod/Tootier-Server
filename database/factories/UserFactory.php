<?php

namespace Database\Factories;

use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Str;
use Ramsey\Uuid\Type\Integer;

class UserFactory extends Factory
{
    protected $model = User::class;
    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            'name_user' => $this->faker->name(),
            'family_user' => $this->faker->name(),
            'role' => 0,
            'username' => $this->faker->userName(),
            'email' => $this->faker->unique()->safeEmail(),
            'first_register_time' => now(),
            'password' => Hash::make("password"), // password
            'longitude_y' => Str::random(10),
            'latitude_x' => Str::random(10),
        ];
    }

}
