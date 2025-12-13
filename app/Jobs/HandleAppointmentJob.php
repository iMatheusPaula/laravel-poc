<?php

namespace App\Jobs;

use App\Enums\AppointmentStatus;
use App\Models\Appointment;
use Carbon\Carbon;
use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;

class HandleAppointmentJob implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    public int $appointmentId;

    /**
     * Create a new job instance.
     */
    public function __construct(int $appointmentId)
    {
        $this->appointmentId = $appointmentId;
    }

    /**
     * Execute the job.
     */
    public function handle(): void
    {
        $appointment = Appointment::find($this->appointmentId);

        // Exit if appointment was deleted or does not exist
        if (!$appointment) {
            return;
        }

        if ($appointment->status !== AppointmentStatus::PENDING) {
            return; // Exit if not pending to handle retries/duplication
        }

        $appointment->update([
            'status' => AppointmentStatus::CONCLUDED,
            'finished_at' => Carbon::now(),
        ]);
    }
}
