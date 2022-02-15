<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Bookmark extends Model
{
    use HasFactory;

    public $timestamps = false;
    protected $table = 't_bookmark';

    protected $fillable = [
        "tooti_id",
        "user_id",
    ];
}
