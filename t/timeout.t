use t;

subtest 'Timeout' => {
    my $res = post-solution :code('sleep 6');

    is $res<Err>, 'Killed for exceeding the 5s timeout.', 'Correct error';

    is floor( $res<Took> / 1e9 ), 5, 'Correct took';
};

subtest 'Timeout with correct output' => {
    my $res = post-solution :code('say "Fizz" x $_ %% 3 ~ "Buzz" x $_ %% 5 || $_ for 1 .. 100; $*OUT.flush; sleep 6;');

    is $res<Err>, 'Killed for exceeding the 5s timeout.', 'Correct error';

    is $res<Out>, slurp('hole/answers/fizz-buzz.txt').trim, 'Correct output';

    nok $res<Pass>, 'Solution fails';
};

done-testing;
