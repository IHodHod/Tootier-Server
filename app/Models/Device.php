<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Device extends Model
{
    use HasFactory;


    public $timestamps = false;
    protected $table = 't_device';

    protected $fillable = [
        "device_uuid",
        "last_time_visit",
        "first_time_register",
        "name_device",
        "mac_address",
        "os",
        "ip",
        "user_id"
    ];
}
