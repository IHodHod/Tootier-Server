<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateBookmark extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('t_bookmark', function (Blueprint $table) {
            $table->id();
            /*
             *
             * TODO
             * Hejazi Pour must do complete right tooti_id
             * */
//            $table->index("tooti_id");
//            $table->foreign('tooti_id')->references('id')->on('t_tooti')->onDelete('cascade');
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
        Schema::dropIfExists('t_bookmark');
    }
}
