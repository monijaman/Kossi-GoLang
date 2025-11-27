#!/usr/bin/perl

my %d2a = ('०'=>'0','१'=>'1','२'=>'2','३'=>'3','४'=>'4','५'=>'5','६'=>'6','७'=>'7','८'=>'8','९'=>'9');

while (<>) {
    my $line = $_;
    # Fix keys by converting Devanagari back to ASCII in the KEY part only
    if ($line =~ /^(\s*)"([^"]*)"(:)/) {
        my ($space, $key, $colon) = ($1, $2, $3);
        foreach my $d (keys %d2a) {
            $key =~ s/$d/$d2a{$d}/g;
        }
        $line =~ s/^(\s*)"[^"]*"(:)/$space"$key"$colon/;
    }
    print $line;
}
