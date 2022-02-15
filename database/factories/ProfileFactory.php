<?php

namespace Database\Factories;

use App\Models\Profile;
use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

class ProfileFactory extends Factory
{
    protected $model = Profile::class;

    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            'image_profile' => "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR-XHbhDY5ZwNVcHm8z3sdNy_CU3HzTgrdIOw&usqp=CAU",
            'image_cover' => "https://wallpaperaccess.com/full/1129092.jpg",
            'biography' => Str::random(40),
            'gender' => 0,
            'birthday' => now(),
            'user_id' => 1,
        ];
    }

}
