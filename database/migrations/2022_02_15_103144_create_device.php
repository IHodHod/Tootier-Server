<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateDevice extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('t_device', function (Blueprint $table) {
            $table->id();
            $table->string("device_uuid")->nullable();
            $table->timestamp("last_time_visit");
            $table->timestamp("first_time_register");
            $table->string("name_device");
            $table->string("mac_address");
            $table->string("os");
            $table->string("ip");
            $table->integer("user_id")->unsigned();
            $table->foreign('user_id')->references('id')->on('t_user')->onDelete('cascade');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('t_device');
    }
}
