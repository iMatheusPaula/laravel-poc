<?php

use App\Http\Controllers\AppointmentController;
use Illuminate\Support\Facades\Route;

Route::controller(AppointmentController::class)->group(function () {
    Route::get('/appointments', 'index');
    Route::post('/appointments', 'store');
    Route::patch('/appointments/{appointment}/cancel', 'cancel');
});
