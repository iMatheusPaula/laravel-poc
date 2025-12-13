<?php

namespace App\Http\Controllers;

use App\Enums\AppointmentStatus;
use App\Models\Appointment;
use Carbon\Carbon;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class AppointmentController extends Controller
{
    public function index(): JsonResponse
    {
        return response()->json(Appointment::query()->latest()->get());
    }

    public function store(Request $request): JsonResponse
    {
        $validated = $request->validate([
            'scheduled_at' => ['required', 'datetime'],
        ]);

        $appointment = Appointment::query()
            ->create([
                'scheduled_at' => $validated['scheduled_at'],
                'status' => AppointmentStatus::PENDING,
            ]);

        return response()->json($appointment, Response::HTTP_CREATED);
    }

    public function cancel(Appointment $appointment): JsonResponse
    {
        if ($appointment->status !== AppointmentStatus::PENDING) {
            return response()->json(['message' => 'Only pending appointments can be canceled.'], Response::HTTP_CONFLICT);
        }

        $appointment->update([
            'status' => AppointmentStatus::CANCELED,
            'canceled_at' => Carbon::now(),
        ]);

        return response()->json($appointment);
    }
}
