<?php

namespace App\Models;

use App\Enums\AppointmentStatus;
use Illuminate\Database\Eloquent\Model;

class Appointment extends Model
{
    protected $fillable = [
        'scheduled_at',
        'status',
        'finished_at',
        'canceled_at'
    ];

    protected function casts(): array
    {
        return [
            'scheduled_at' => 'datetime',
            'status' => AppointmentStatus::class,
            'finished_at' => 'datetime',
            'canceled_at' => 'datetime'
        ];
    }
}
