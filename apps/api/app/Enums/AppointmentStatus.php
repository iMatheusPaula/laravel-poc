<?php

namespace App\Enums;

enum AppointmentStatus: string
{
    case PENDING = 'pending';
    case CONCLUDED = 'concluded';
    case CANCELED = 'canceled';
}
