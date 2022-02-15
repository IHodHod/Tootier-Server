<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateUsersTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('t_user', function (Blueprint $table) {
            $table->id();
            $table->string('username' , 128)->unique();
            $table->string('email' , 128)->unique();
            $table->string('name_user' , 50);
            $table->string('family_user' , 50);
            $table->string("role");
            $table->string('password');
            $table->string("latitude_x");
            $table->string("longitude_y");
            $table->timestamp('first_register_time');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('t_user');
    }
}
