<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateProfile extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('t_profile', function (Blueprint $table) {
            $table->id();
            $table->string("image_profile");
            $table->string("image_cover");
            $table->string("biography");
            $table->integer("gender");
            $table->timestamp("birthday");
            $table->integer("user_id")->unsigned();
            $table->foreign("user_id")->references("id")->on("t_user")->onDelete("cascade");
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('t_profile');
    }
}
