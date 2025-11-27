#!/usr/bin/perl

my %d2a = ('०'=>'0','१'=>'1','२'=>'2','३'=>'3','४'=>'4','५'=>'5','६'=>'6','७'=>'7','८'=>'8','९'=>'9');

my $in_specs = 0;
while (<>) {
    my $line = $_;
    
    if ($line =~ /specs\s*:=\s*map/) {
        $in_specs = 1;
    } elsif ($in_specs && $line =~ /^\s*\}/) {
        $in_specs = 0;
    }
    
    # Convert all Devanagari to ASCII in specs section
    if ($in_specs) {
        foreach my $d (keys %d2a) {
            $line =~ s/$d/$d2a{$d}/g;
        }
    }
    
    print $line;
}
