<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;

class LintProject extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'lint {--fix : Fix errors that can be fixed using phpcbf}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Lint project files';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return int
     */
    public function handle()
    {
        if ($this->option("fix")) {
            exec("./vendor/bin/phpcbf ./app ./database ./routes ./tests ./config", $output);
            $this->comment(implode(PHP_EOL, $output));
            return 0;
        }

        exec("./vendor/bin/phpcs ./app ./database ./routes ./tests ./config", $output);
        $this->comment(implode(PHP_EOL, $output));
        return 0;
    }
}
