<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Profile extends Model
{
    use HasFactory;

    public $timestamps = false;
    protected $table = 't_profile';

    protected $fillable = [
        "image_profile",
        "image_cover",
        "biography",
        "gender",
        "birthday",
        "user_id"
    ];
}
