<?php

namespace Database\Factories;

use App\Models\Device;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

class DeviceFactory extends Factory
{
    protected $model = Device::class;

    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            'device_uuid' => $this->faker->uuid(),
            'last_time_visit' => now(),
            'first_time_register' => now(),
            'name_device' => "Samsung " . Str::random(5),
            'mac_address' => $this->faker->macAddress(),
            'user_id' => 1,
            'os' => 'Android os ',
            'ip' => $this->faker->ipv4(),
        ];
    }

}
